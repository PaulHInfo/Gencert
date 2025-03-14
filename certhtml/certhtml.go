package certhtml

import (
	"fmt"
	"gencert/cert"
	"os"
	"path"
	"text/template"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outputDir string) (*HtmlSaver, error) {
	var h *HtmlSaver
	/*err := os.Mkdir(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}*/
	h = &HtmlSaver{OutputDir: outputDir}
	return h, nil
}

func (p *HtmlSaver) Save(cert cert.Cert) error {
	t, err := template.New("Certificate").Parse(tpl)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%v.html", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, cert)

	if err != nil {
		return err
	}
	fmt.Println("HTML SAVED")
	return nil
}

var tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
        <title>{{.LabelTitle}}</title>
        <style>
            body {
                text-align: center;
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
            }

            h1 {
                font-size: 3em;
            }

        </style>
    </head>
    <body>
        <h1>{{.LabelCompletion}}</h1>
        <h2><em>{{.LabelPresented}}</em></h2>
        <h1>{{.Name}}</h1>
        <h2>{{.LabelParticipation}}</h2>
        <p>
            <em>{{.LabelDate}}</em>
        </p>
    </body>
</html>


`
