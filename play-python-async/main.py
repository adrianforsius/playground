from pprint import pprint as p
import asyncio

import grequests
import requests


async def main():
    # group = await asyncio.gather(*[crawl(), crawl(), crawl(), crawl()], return_exceptions=True,)
    requests.get("https://www.google.com")


async def crawl():
    result = []
    for i in range(1, 5):
        r = await grequests.get("https://www.google.com")
        result.append(r)
    return result


if __name__ == "__main__":
    # asyncio.run(main(), debug=True)
    requests.get("https://www.google.com")