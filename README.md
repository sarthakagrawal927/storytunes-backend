# Schema Design


```js

Story: {
    heading
    likes_count
    collaborator_count
    tags
    story_preview : 1st 3 sentences = (or AI summarizer)
    authorUsername
    story_image
    status
}

Sentence {
    storyId
    position
    rank
    status : {
        isWinnerByVote
        rejectedByAuthor
        isWinnerByAuthor
        locked
    }
    likes_count
    text
    imageUrl
}

// mongoDB
// [
//     sentence1 : {
//         active :{
//              text, author
//         },
//         past: [
//             {
//                 text, author, pts
//             }
//         ]
//     }
// ]

/*
Ram is shyam - U1
sarhtak is not handsome - U1, U2
asdasd asd  adqw - U1
adasd qwd wd qwdqw - U1
*/

1 1 U1
2 1 U2
2 2 U1
3 1 U1
3 1 U1

1 Story - 50 sentence - 3 alternative -> 150 row - 450 row

User_Story {
    storyId
    userId
    status // bookmark or creator
}

/*
 b c
 0 0 - 0
 0 1 - 1
 1 0 - 2
 1 1 - 3
*/

EntityLikes{ // 10million rows
    entityType - (Sentence or Post)
    sentenceID
    storyID
    userID
}

User {
    firstName
    lastName
    email
    username
}
```


select * from entity_likes from etype = 1 and postId = 2;


