package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Antonboom/testifylint/internal/checkers"
)

var (
	baseTestsGenerator = BaseTestsGenerator{}
	baseTestsPath      = filepath.Join("analyzer", "testdata", "src", "base-tests", "base_test.go") // TODO: common code with checkers
)

var generators = []CheckerTestsGenerator{
	BoolCompareCasesGenerator{},
	// ComparesCasesGenerator{},
	// EmptyCasesGenerator{},
	// ErrorCasesGenerator{},
	// ErrorIsCasesGenerator{},
	// ExpectedActualCasesGenerator{},
	// FloatCompareCasesGenerator{},
	// LenCasesGenerator{},
	// RequireErrorCasesGenerator{},
	// SuiteNoExtraAssertCallCasesGenerator{},
}

func init() {
	genForChecker := make(map[string]struct{}, len(generators))
	for _, g := range generators {
		name := g.CheckerName()
		if name == "" {
			panic(fmt.Sprintf("%T: checker name not defined", g))
		}
		if _, ok := genForChecker[name]; ok {
			panic(fmt.Sprintf("several generators for checker %q", name))
		}
		genForChecker[name] = struct{}{}
	}

	for _, ch := range checkers.All() {
		if _, ok := genForChecker[ch]; !ok {
			log.Printf("[WARN] No generated tests for %q checker\n", ch)
		}
	}
}

func main() {
	if err := generateBaseTests(); err != nil {
		log.Panic(err)
	}

	if err := generateCheckersTests(); err != nil {
		log.Panic(err)
	}
}

func generateBaseTests() error {
	return genTestFilesPair(baseTestsGenerator, baseTestsPath)
}

func generateCheckersTests() error { // TODO: use genTestFilesPair
	for _, g := range generators {
		output := toCheckersPath(g.CheckerName(), strings.ReplaceAll(g.CheckerName(), "-", "_")+"_test.go")

		dir := filepath.Dir(output)
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("rm tests dir: %v", err)
		}
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("mkdir tests dir: %v", err)
		}

		tmplData := g.Data()

		if err := genGoFileFromTmpl(output, g.ErroredTemplate(), tmplData); err != nil {
			return fmt.Errorf("generate test file: %v", err)
		}

		if goldenTmpl := g.GoldenTemplate(); goldenTmpl != nil { // TODO: in such case?
			if err := genGoFileFromTmpl(output+".golden", goldenTmpl, tmplData); err != nil {
				return fmt.Errorf("generate golden file: %v", err)
			}
		}
	}
	return nil
}

func toCheckersPath(dirsFile ...string) string {
	return filepath.Join(append([]string{"analyzer", "testdata", "src", "checkers"}, dirsFile...)...)
}

func genTestFilesPair(g TestsGenerator, path string) error {
	tmplData := g.Data()

	if err := genGoFileFromTmpl(path, g.ErroredTemplate(), tmplData); err != nil {
		return fmt.Errorf("generate test file: %v", err)
	}

	if err := genGoFileFromTmpl(path+".golden", g.GoldenTemplate(), tmplData); err != nil {
		return fmt.Errorf("generate golden file: %v", err)
	}
	return nil
}

func genGoFileFromTmpl(output string, tmpl *template.Template, data any) error {
	b := bytes.NewBuffer(nil)
	if err := tmpl.Execute(b, data); err != nil {
		return fmt.Errorf("execute cases tmpl: %v", err)
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		_ = os.WriteFile(output, b.Bytes(), 0o644) // For debug.
		return fmt.Errorf("format %s: %v", output, err)
	}

	return os.WriteFile(output, formatted, 0o644)
}
