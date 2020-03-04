import os
import json
import datetime
from feedparser import parse
from flask import Blueprint, Response, request
from flask_restful import Resource, Api
from bs4 import BeautifulSoup
import pprint


def image_from_feed(entry):
    if hasattr(entry, "content"):
        post = entry.content.pop().value
    elif hasattr(entry, "summary_detail"):
        post = entry.summary_detail.value
    else:
        post = entry.summary
    soup = BeautifulSoup(post, 'html.parser')
    try:
        return soup.find('img')['src']
    except TypeError:
        return None


rss_urls = ['http://www.vsnp.net/index.rdf',
            'http://blog.livedoor.jp/dqnplus/index.rdf']


app = Blueprint('fetch_latest_articles', __name__)
api = Api(app)


class RSSReaderResource(Resource):
    def get(self):
        articles_json = []
        for rss_url in rss_urls:
            author = parse(rss_url).feed.title
            entries = parse(rss_url).entries
            articles_json.extend([{'title': entry['title'], 'author': author,
                                   'url': entry['link'], 'urlToImage': image_from_feed(entry), 'publishedAt': 'tbd'} for entry in entries])
            resp = {'articles': articles_json}
        return Response(response=json.dumps(resp), status=200)


api.add_resource(RSSReaderResource, '/fetch_latest_articles')
# pprint.pprint(vs)
