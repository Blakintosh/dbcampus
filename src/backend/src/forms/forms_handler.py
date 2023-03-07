from create_form import createHandler
from retrieve_results import retrieveHandler
import json
import sys


def main(data):
    try:
        if (data["function"] == "create"):
            return createHandler(data)
        elif (data["function"] == "retrieve"):
            # pass
            return retrieveHandler(data)
        else:
            print("bad json")
            return False
    except:
        print("bad json")
        return False



if __name__ == '__main__':
    data = sys.argv.pop()
    main(json.loads(data))



