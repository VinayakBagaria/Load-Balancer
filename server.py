from flask import Flask

app = Flask('Server-1')


@app.route('/')
def index():
    return app.name


app.run()
