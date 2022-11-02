People can have multiple tags.
Projects should only have one tag

looks something like:
```
{
    "name": "",
    "gitlab": "",
    "mastodon": "",
    "github": "",
    "reddit": "",
    "matrix": "",
    "twitter": "",
    "youtube": "",
    "website": "",
    "donate": "",
    "avatar": "",
    "description": "",
    "tags": [

    ]
}
```

example:

```
{
    "name": "Seth Simmons",
    "github": "sethforprivacy",
    "reddit": "fort3hlulz",
    "matrix": "sethsimmons:monero.social",
    "twitter": "sethforprivacy",
    "youtube": "OptOutPodcast",
    "website": "https://sethforprivacy.com/",
    "donate": "86JzKKyZvtEC98y6zJxCCVfcA3r75XngPBjpYDE6zRR36keNGMHwZomDjMCv1oCYB2j9myiFqEJQF3JtnhKdfX546T91eaY",
    "avatar": "https://avatars.githubusercontent.com/u/40500387?v=4",
    "description": "I like Monero [getmonero.org](www.getmonero.org)",
    "tags": [

        "OptOut",
        "Events"
    ]
}
```

the corresponding 'OptOut' project:
```
{
    "name": "OptOut Podcast",
    "github": "",
    "twitter": "optoutpod",
    "donate": "https://www.optoutpod.com/support/",
    "type": "project",
    "website": "https://www.optoutpod.com/",
    "youtube": "OptOutPodcast",
    "avatar": "",
    "description": "Welcome to Opt Out, where I sit down with passionate people to learn why privacy matters to them, the tools and techniques they\u2019ve found and leveraged, and where we encourage and inspire others towards personal privacy and data-sovereignty.",
    "tags": [
        "OptOut"
    ]
}
```

## I want to be on the list
[how to make a PR on github](https://docs.github.com/en/github/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request)

Lets say my name is 'John Doe'. I am a member of the 'Monero Policy Workgroup' and i want to be added on the devlist website.
I fork this repo, then i create a new file in 'people' called (name does not matter) johndoe.json. The 'tag' for the MPW is, "MPW" so i need to add that as a tag.
```
~ Contents of people/johndoe.json ~
{
    "name": "John Doe",
    "avatar": "./images/people/johndoe.png" <- If empty, default icon will be used. You can either add a link or store your image file in /static/images/people.
    "github": "", <- Just provide the username for this and reddit/matrix/gitlab
    "reddit": "",
    "matrix": "",
    "twitter": "",
    "youtube": "", <- The /c/ channel name (not the uid or the link wont work)
    "website": "",
    "donate": "", <- A web url can be used to take people to your page OR put an address/openalias here to display on the site
    "avatar": "", <- If you leave this empty, your twitter / or github / of default icon will be auto filled in because i like saving people time.
    "description": "", <- this will default to "I like monero getmonero.org" if left empty
    "tags": [
        "MPW" <- We've added the tag
    ]
}

```

Then submit a PR and wait for the 'human' (probably me) to review it.

John Doe now wants to create a new project! called 'cyberpunks'! so he needs to create a file in the projects folder, using the same template:

```
~ Contents of projects/cyberpunks.json~
{
    "name": "Cyber Punks",
    "github": "",
    "reddit": "",
    "matrix": "",
    "twitter": "",
    "youtube": "",
    "website": "",
    "donate": "",
    "avatar": "",
    "description": "We are a team of CyberPunks!",
    "tags": [

        "Punk" <- anyone who wants to be shown in this group can add a PR with "Punk" added to their tags.
    ]
}
```

John Doe can then , edit his people/johndoe.json and add the "Punk" to his tags as shown below:
```
~ Contents of people/johndoe.json ~
{
    "name": "John Doe",
    "github": "",
    "reddit": "",
    "matrix": "",
    "twitter": "",
    "youtube": "",
    "website": "",
    "donate": "",
    "avatar": "",
    "description": "",
    "tags": [
        "MPW",
        "Punks" <- We've added the new tag (notice no comma after the last tag)
    ]
}
```
