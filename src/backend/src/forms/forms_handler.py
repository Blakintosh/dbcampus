from create_form import createHandler
from retrieve_results import retrieveHandler
import json
import sys


def main(data):
    try:
        if (data["function"] == "create"):
            
            res = createHandler(data)
            print(res)
            
            #converting to json string
            json_object = json.dumps(res, indent=4)
            
            # Writing to result json
            with open("res.json", "w") as outfile:
                outfile.write(json_object)

        elif (data["function"] == "retrieve"):
            
            res = retrieveHandler(data)
            print(res)
            
            #converting to json string
            json_object = json.dumps(res, indent=4)
            
            # Writing to result json
            with open("res.json", "w") as outfile:
                outfile.write(json_object)
            
        else:
            print("bad json")
            return False
        
    except:
        print("bad json")
        return False



if __name__ == '__main__':
    data = sys.argv.pop()
    main(json.loads(data))



