package structs

import "fmt"

type MetaData struct {
	CreatedBy string
	CreatedAt string
}

type Post struct {
	MetaData    // Embedding, Ek olarak buna bir isim versedik cagırırken kesinlinle o ismi kullanmak zorunda kalacaktık
	Title       string
	ContentBody string
}

func main() {
	post := Post{
		MetaData: MetaData{
			CreatedBy: "Ryuka",
			CreatedAt: "today",
		},
		Title:       "Ucak Tekerlekleri",
		ContentBody: "Bence oe",
	}

	// post.MetaData.CreatedBy ve post.CreatedAt post.CreatedBy farkı :
	//	Hiçbir fark yoktur. İkisi de RAM'de aynı adrese gider

	// ek olarak burda formatlama için ek olarak kullanım onerisi : ??
	fmt.Printf("\nYazar : [%v] - Konu : %v \n", post.CreatedBy, post.Title)

	fmt.Printf("\nPost Detayları : %+v", post)
}
