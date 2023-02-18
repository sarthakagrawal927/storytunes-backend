// Story Structure in Redis

{
    contributors: {
        id: "username"
    }
    upvotes: {
        id: "username"
    }
    summary: "text"
    lines: {
        idx: {
            primary: "text"
            upvotes: {
                id: "username"
            }
            alternates: {
                rank_status: {
                    text:"text"
                    upvotes: {
                        id: "username"
                    }
                }
            }
        }
    }
}