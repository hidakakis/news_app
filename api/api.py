import os
from flask import Flask
from flask_cors import CORS
from controller import rss_reader
from controller import fetch_latest_articles


app = Flask(__name__)
CORS(app)
app.register_blueprint(fetch_latest_articles.app, url_prefix='/api')


if __name__ == '__main__':
    app.run(host='0.0.0.0')
