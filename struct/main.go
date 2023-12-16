package main

import "fmt"

type Movie struct{
    title string
    mType string
    rating int
    Comments
}

type Comments struct{
    author string
    body string
}

func main(){
    fmt.Println("Structs")

    film := Movie{
        title: "All About Eve",
        mType: "Drama",
        rating: 9,
        Comments: Comments{
            author: "Joseph Mankiewicz",
            body: "The film stars Bette Davis as Margo Channing, a highly regarded but aging Broadway star, and Anne Baxter as Eve Harrington, an ambitious young fan who maneuvers herself into Channing's life, ultimately threatening Channing's career and her personal relationships. The film co-stars George Sanders, Celeste Holm, Gary Merrill, and Hugh Marlowe, and features Thelma Ritter, Marilyn Monroe in one of her earliest roles, Gregory Ratoff, Barbara Bates and Walter Hampden.",
        },
    }
    fmt.Printf("Name: %v\nType: %v\nRating: %v\nAuthor:%v\nAbout: %v", film.title, film.mType, film.rating, film.author, film.body)
}
