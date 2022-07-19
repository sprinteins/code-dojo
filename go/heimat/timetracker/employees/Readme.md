# Go Heimat


```txt
┌──────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                                          │
│                                                                                                          │
│                                                     O                                                    │
│                                                    /|\                                                   │
│                                                    / \                                                   │
│                                                                                                          │
│                                                   USER                                                   │
│                                                     │                                                    │
│                                                     │                                                    │
│                     /api/employees                  │                     /api/projects                  │
│                      ┌──────────────────────────────┼─────────────────────────────┐                      │
│                      │                      /api/timetracker                      │                      │
│                      │                              │                             │                      │
│                      ●                              ●                             ●                      │
│                      ◡                              ◡                             ◡                      │
│                      │                              │                             │                      │
│                    :5040                          :5050                         :5060                    │
│                      │                              │                             │                      │
│                      │                              │                             │                      │
│       ┌──────────────┴─────────────┐ ┌──────────────┴─────────────┐ ┌─────────────┴────────────┐         │
│       │          Employees         │ │         TimeTracker        │ │         Projects         │         │
│       ├────────────────────────────┤ ├────────────────────────────┤ ├──────────────────────────┤         │
│       │                            │ │                            │ │                          │         │
│       │+ Add(                      │ │+ AddTime(u,p,t Minute)     │ │+ AddProject(name string) │         │
│       │    firstName,              │ │+ WorkedTime(u) []Minute    │ │+ AddUser(                │         │
│       │    lastName,               │ │+ WorkedTimeByProject(      │ │    p string,             │         │
│       │    username,               │ │    u,                      │ │    u string,             │         │
│       │)                           │ │    p,                      │ │  )                       │         │
│       │+ Activate(u)               │ │  ) Minute                  │ │+ RemoveUser(             │         │
│       │+ Deactivate(u)             │ │                            │ │    p string,             │         │
│       │+ IsActive(u) Bool          │ │                            │ │    u string,             │         │
│       │                            │ │                            │ │  )                       │         │
│       │                            │ │                            │ │+ Projects() []Project    │         │
│       │                            │ │                            │ │+ ProjectsByUser(u string)│         │
│       └────────────────────────────┘ └────────────────────────────┘ └──────────────────────────┘         │
│                                                                                                          │
│                                                                                                          │
├──────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│                                        ┏━━━━━━━━━━━━━━━━━━━━━━━━┓                      ┌───────────────┐ │
│ Goofy · Architecture · 08.07.2020      ┃       TOP SECRET       ┃                      │      SE       │ │
│                                        ┗━━━━━━━━━━━━━━━━━━━━━━━━┛                      └───────────────┘ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```


> Every services should:
> - [ ] return at least one HTTP Status error
> - [ ] persist data
> - [ ] HTTP middleware

## Employees

- `Add(): void`
  1. checks if user already exists (return error: "user exists")
  2. hashes password
  3. persists new user
- `pwdHash(usr): string`
    1. returns the password hast for a given user
    2. return error if user does not exists

## Transaction

- `send(userA, userB, amount): void`
    1. checks if user token is valid (external call to `Auth`
    2. checks if user exists (external call to `Account`)
    3. checks if `fromUser` is the same as logged in user
    4. checks if `toUser` exists (external call to `Account`)
    5. checks if `fromUser` has enough capital
    6. records transaction
       1. (optional) recalculates assets for each user
- `assets(userA): int`
    1. checks if token is valid
    2. checks if `usr` is the same as logged in user
    3. checks if user exists
    4. calculate assets based on transactions 
       1. (alternative) returns pre-calculated value
- `deposit(usr, amount)`
    1. checks if token is isValid
    2. checks if user exists
    3. checks if `usr` is the logged in user
    4. records transaction 
       1. (optional) recalculates assets for user

## Auth
- `login(usr, pwd): Token`
    1. checks if user exists
    2. retrieves `pwdHash` (external call)
    3. calculates password hash based on given `pwd`
    4. checks if calculated password hash and `pwdHash` are the same
    5. if yes return a token 
       1. if no returns an error
    6. records failed attempts
- `isValid(token): Boolean`
    1. checks if token signature is valid
    2. returns error if not
- `isLoggedInUser(usr, token): Boolean`
    1. checks if token is valid (return error if not)
    2. return true if the user in the token is the same, otherwise false