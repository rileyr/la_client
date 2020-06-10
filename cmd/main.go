package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/rileyr/littleaspen"
)

func main() {
	fmt.Println("starting")

	key := os.Getenv("LITTLE_ASPEN_API_KEY")
	if key == "" {
		panic("missing LITTLE_ASPEN_API_KEY")
	}
	c := littleaspen.New(key)

	documents, err := c.GetDocuments()
	if err != nil {
		panic(err)
	}

	for _, doc := range documents {
		spew.Dump(doc)
		versions, err := c.GetContentVersions(doc.Slug)
		if err != nil {
			panic(err)
		}

		for _, v := range versions {
			vers, err := c.GetContentVersion(doc.Slug, v.Slug)
			if err != nil {
				panic(err)
			}
			spew.Dump(vers)

			accs, err := c.GetAcceptances(doc.Slug, v.Slug)
			if err != nil {
				panic(err)
			}

			if len(accs) > 0 {
				for _, acc := range accs {
					a, err := c.GetAcceptance(doc.Slug, v.Slug, acc.ExternalID)
					if err != nil {
						panic(err)
					}
					spew.Dump(a)
				}
			} else {
				uid := uuid.New()
				a, err := c.CreateAcceptance(doc.Slug, v.Slug, uid.String(), nil)
				if err != nil {
					panic(err)
				}
				spew.Dump(a)
			}
		}
	}

	fmt.Println("done")
}
