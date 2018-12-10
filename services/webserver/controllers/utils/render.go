package utils

import (
	"os"
	"log"
	"net/http"
	"path/filepath"
	"html/template"
    "github.com/smart-evolution/smarthome/utils"
	"github.com/smart-evolution/smarthome/models/agent"
    "github.com/smart-evolution/smarthome/models/page"
	"github.com/coda-it/gowebserver/session"
)


// RenderTemplate - helper for page rendering
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, sm session.ISessionManager) {
    sessionID, _ := GetSessionID(r)
    isLogged := sm.IsExist(sessionID)
    isPrivate := IsRequestFromIntranet(r)

    if !isLogged {
        ClearSession(w)

        if r.URL.Path != "/login" && r.URL.Path != "/login/register" {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
        }
    }

    if !isPrivate && r.URL.Path == "/login/register" {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }

    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

    if err != nil {
        log.Fatal(err)
    }

    menu := make([]agent.Agent, 0)
    for _, a := range agent.Agents {
        menu = append(menu, a)
    }

    params := make(map[string]interface{})
    params["menu"] = menu

    templateModel := page.Page{
        Version: utils.VERSION,
        Title: name,
        IsLogged: isLogged,
        Params: params,
    }

	template := template.Must(
        template.ParseFiles(
            dir + "/views/" + name + ".html",
            dir + "/views/navigation.html",
            dir + "/views/view.html",
        ),
    )
	template.ExecuteTemplate(w, "base", templateModel)
}