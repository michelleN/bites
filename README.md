# bites

small CRUD API written in Go with the Spin Framework

## Tutorial

1. Create a new repo on GitHub. Grab a License, gitignore file for Go and anything else you might want.
2. git clone that new repo onto your machine and cd into the bites directory
3. spin new bites --init. You'll get some prompts. You can hit enter to accept defaults or fill in whatever you desire.
4. Inspect your spin.toml file. You'll see you have one component called bites and it gets triggered by any path.
5. Test by running `$ spin build` in your terminal. If you run into errors they might be because you don't have `tinygo` installed. Make sure you install that and run `spin build` again. Then, run `$ spin up`. You'll see that the app is being served on http://127.0.0.1:3000 and if you curl that address you'll get some text: "Hello Fermyon!".
6. Deep dive into Go code
If you go to main.go, you'll see an empty main function and an init function that handles the request using the Handle function from the spinhttp SDK that has been imported. With Spin, your go files need to have an empty main function and your logic needs to be in the init function. Here you can change the content if you like, run `spin build && spin up` to see your changes.
7. Now what we really want is a way to create/read/delete/list all bites so let's start with creating a bite. Let's do that by properly handling POST requests to our app.

In the spin Go SDK, there is a router you can use to easily handle requests with different Methods. Create a router by calling NewRouter(): `router := spinhttp.NewRouter()`. Then,
set how you want to handle POST requests like this:

```go
router.POST(<route>, <some-function-that-handles-the-request)

```

8. Now test using `spin up --build` and `curl`:

```bash
curl -i  -X POST localhost:3000/helloworld
```

9. Now, let's actually store the bite somewhere. Let's use SQLite. You've probably used SQLite before. There is nothing to install and Spin will configure everything for you when you `spin up`. All you need to do add a few lines to the spin.toml file, create the table structure you want to use and use the spin sqlite SDK in your code

In `spin.toml` file add:

```toml
[component.bites]
sqlite_databases = ["default"]
```

In terminal, in directory:

```bash
$ spin up --sqlite "CREATE TABLE IF NOT EXISTS bites (id INTEGER PRIMARY KEY AUTOINCREMENT, content TEXT NOT NULL)"

Logging component stdio to ".spin/logs/"
Storing default SQLite data to ".spin/sqlite_db.db"

Serving http://127.0.0.1:3000
Available Routes:
  bites: http://127.0.0.1:3000 (wildcard)
```

_Note: If you make a mistake setting up your sqlite tables, just delete the sqlite file(s) in the .spin directory or you can rm -rf the entire .spin directory._

In code:

```go
import (
    "github.com/fermyon/spin/sdk/go/sqlite"
)

func init(){
    ...
    sqlite.Open("default")

}
```
