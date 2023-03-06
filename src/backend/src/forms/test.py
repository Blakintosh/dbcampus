#!/usr/local/bin/python
import sys

def main(data):
        print("out "+data)
        return data

if __name__ == '__main__':
    data = sys.argv.pop()
    main(data) 
