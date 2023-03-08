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
            res = retrieveHandler(data)
            print(res)
            
            #Serializing json
            json_object = json.dumps(res, indent=4)
            
            # Writing to sample.json
            with open("res.txt", "w") as outfile:
                outfile.write(json_object)
            # return json.dumps(res)
        else:
            print("bad json")
            return False
    except:
        print("bad json")
        return False



if __name__ == '__main__':
    data = sys.argv.pop()
    main(json.loads(data))



