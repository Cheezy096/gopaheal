package gopaheal

func GetPosts(tags []string) (allPosts []string, err error) {
	url := constructUrl(1, tags)

	count, err := getLastPage(url)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= count; i++ {
		// Range over all pages

		url := constructUrl(i, tags)
		posts, err := getPosts(url)
		if err != nil {
			return nil, err
		}

		allPosts = append(allPosts, posts...)
	}

	return allPosts, nil
}

func GetPostsSlice(tags []string) (allPostsSliced [][]string, err error) {
	url := constructUrl(1, tags)

	count, err := getLastPage(url)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= count; i++ {
		// Range over all pages

		url := constructUrl(i, tags)
		posts, err := getPosts(url)
		if err != nil {
			return nil, err
		}

		allPostsSliced = append(allPostsSliced, posts)

	}

	return allPostsSliced, nil
}
