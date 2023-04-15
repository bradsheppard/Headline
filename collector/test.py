import json
from typing import List
from duckduckgo_search import ddg_news

results: List[str] = ddg_news('Metallica')

for result in results:
    print(json.dumps(result, indent=4))

print(results)
