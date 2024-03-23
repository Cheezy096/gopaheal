package gopaheal

import (
	"fmt"
	"os"
	"testing"
)

var testTag = []string{"Higurashi_When_They_Cry"}

func TestNormal(t *testing.T) {
	_, err := GetPosts(testTag)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSlice(t *testing.T) {
	_, err := GetPostsSlice(testTag)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteNormal(t *testing.T) {
	postsSlice, err := GetPosts(testTag)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create(fmt.Sprintf("cache_%s.txt", testTag[0]))
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	file.Truncate(0)

	for i := range postsSlice {
		_, err = file.WriteString(fmt.Sprintf("%s\r\n", postsSlice[i]))
		if err != nil {
			t.Fatal(err)
		}
	}

	file.Sync()
}

func TestWriteSlice(t *testing.T) {
	postsSlice, err := GetPostsSlice(testTag)
	if err != nil {
		t.Fatal(err)
	}

	for i := range postsSlice {
		file, err := os.Create(fmt.Sprintf("./slice/%d_cache_%s.txt", i+1, testTag[0]))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		file.Truncate(0)

		for ii := range postsSlice[i] {
			_, err = file.WriteString(fmt.Sprintf("%s\r\n", postsSlice[i][ii]))
			if err != nil {
				t.Fatal(err)
			}
		}

		_, err = file.WriteString("\r\n")
		if err != nil {
			t.Fatal(err)
		}

		file.Sync()
	}
}
