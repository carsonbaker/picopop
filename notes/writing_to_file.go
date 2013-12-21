  // write out the file to pcm for debugging

  fo, err := os.Create("output.pcm")
  if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
      if err := fo.Close(); err != nil {
          panic(err)
      }
  }()

  if _, err := fo.Write(buffer); err != nil {
    panic(err)
  }

  fmt.Println("Done.")
