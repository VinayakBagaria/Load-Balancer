import sys
from flask import Flask

app = Flask(sys.argv[1])


@app.route('/')
def index():
    return app.name


app.run(port=sys.argv[2])
