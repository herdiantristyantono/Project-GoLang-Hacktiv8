package user

import (
	"log"
	"net/http"
)

type Articles struct {
	ArticleId	int32		`json:"article_id"`
	UserId 		int32		`json:"user_id"`
	Contents    string		`json:"contents"`
	Published	bool		`json:"published"`
}

func (m *Module) HomeUser(w http.ResponseWriter, r *http.Request) {
	rows, err := m.Queries.SelectHome.Query()
	if err != nil {
		log.Println("Failed to insert data")
		return
	}

	var rowsScanArr []Articles

	//Fect data to struct
	for rows.Next() {
		var rowsScan = Articles{}

		err := rows.Scan(&rowsScan.ArticleId, &rowsScan.UserId, &rowsScan.Contents, &rowsScan.Published)
		if err != nil {
			return
		}

		// Append for ervery next row
		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	err = m.Template.ExecuteTemplate(w, "home.html", rowsScanArr)
	if err != nil {
		log.Println(`error execute template login, err : `, err)
		return
	}
}
