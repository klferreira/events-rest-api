db.events.insertMany([
    { 
        name: "Random Party", 
        sessions: [
            ISODate("2019-12-19T20:12:30.348Z"), 
            ISODate("2020-01-15T20:30:16.348Z")
        ], 
        place: "Random Venue", 
        tags: ["electronic", "dance"],
        interested: 0,
    },
    { 
        name: "AC/DC Live Performance", 
        sessions: [
            ISODate("2020-03-20T20:12:30.348Z"), 
            ISODate("2020-04-20T20:30:16.348Z")
        ], 
        place: "Another Random Venue", 
        tags: ["rock", "classic-rock", "live"],
        interested: 0,
    },
]);