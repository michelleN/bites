# bites

small CRUD API written in Go

## Tutorial

1. Create a new repo on GitHub. Grab a License, gitignore file for Go and anything else you might want.
2. git clone that new repo onto your machine and cd into the bites directory
3. spin new bites --init. You'll get some prompts. You can hit enter to accept defaults or fill in whatever you desire.
4. Inspect your spin.toml file. You'll see you have one component called bites and it gets triggered by any path.
5. Test by running `$ spin build` in your terminal. If you run into errors they might be because you don't have `tinygo` installed. Make sure you install that and run `spin build` again. Then, run `$ spin up`. You'll see that the app is being served on http://127.0.0.1:3000 and if you curl that address you'll get some text: "Hello Fermyon!".
6. Deep dive into Go code
If you go to main.go, you'll see an empty main function and an init function that handles the request using the Handle function from the spinhttp SDK that has been imported. With Spin, your go files need to have an empty main function and your logic needs to be in the init function. Here you can change the content if you like, run `spin build && spin up` to see your changes.
7. Now what we really want is a way to create/read/delete/list all bites so let's start with creating a bite. Let's do that by properly handling POST requests to our app.