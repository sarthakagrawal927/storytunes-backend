# Timeline

## POC

- Basic Website
  - Authentication (Next-Auth)
  - Integrate GIPHY

- Backend
  - Realtime (ELIXIR)
    - Feed (Collaborators & Likes on each Story)
    - Single Story (New sentence, new vote, active collaborators)
  - Basic Routes (GO):
    - GET
      - /stories
      - /stories/:ID
      - /user

    - Post
      - /vote/:storyID/:sentenceID
      - /story
      - /comment