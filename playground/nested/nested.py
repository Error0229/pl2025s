def greet(name):
    greeting = "Hey, "

    def who(name):
        return greeting + name + "!"
    return who(name)
