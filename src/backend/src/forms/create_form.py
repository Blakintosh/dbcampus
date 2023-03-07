from __future__ import print_function

from apiclient import discovery
from httplib2 import Http
from oauth2client import client, file, tools
import sys
import json

def createHandler(data):
    form_service = createService()
    form = createForm(form_service, "test form")
    for x in data["questions"]:
        
        if x["type"] == "scale":
            addQuestion(form_service, form, createScaleQuestion(x["title"], x["questionID"]))
        elif x["type"] == "choice":
            addQuestion(form_service, form, createYNQuestion(x["title"], x["questionID"]))

def createService():
    SCOPES = "https://www.googleapis.com/auth/forms.body"
    DISCOVERY_DOC = "https://forms.googleapis.com/$discovery/rest?version=v1"

    store = file.Storage('token.json')
    creds = None
    if not creds or creds.invalid:
        flow = client.flow_from_clientsecrets('credentials.json', SCOPES)
        creds = tools.run_flow(flow, store)

    form_service = discovery.build('forms', 'v1', http=creds.authorize(
        Http()), discoveryServiceUrl=DISCOVERY_DOC, static_discovery=False)
    
    return form_service


def createForm(form_service, title):

    NEW_FORM = {
        "info": {
            "title": title,
        }
    }
    form = form_service.forms().create(body=NEW_FORM).execute()
    return form



def createScaleQuestion(title, qID, low = 1, high = 5, lowLabel = "Strongly Disagree", highLabel = "Strongly Agree"):
    NEW_QUESTION = {
        "requests": [{
            "createItem": {
                "item": {
                    'title': title,
                    'questionItem': {
                        'question': {
                            "questionId": qID,
                            'required': True,
                            'scaleQuestion': {
                                'low': low,
                                'high': high,
                                'lowLabel': lowLabel,
                                'highLabel': highLabel
                            }
                        }
                    }
                },
                "location": {
                    "index": 0
                }
            }
        }]
    }
    return NEW_QUESTION

def createYNQuestion(title):
    NEW_QUESTION = {
        "requests": [{
            "createItem": {
                "item": {
                    'title': title,
                    'questionItem': {
                            # 'questionId': title,
                        'question': {
                            'required': True,
                            'choiceQuestion': {
                                'type': 'RADIO',
                                'options': [{'value': 'yes'}, 
                                            {'value': 'no'}]
                            }
                        }
                    }     
                },
                "location": {
                    "index": 0
                }
            }
        }]
    }
    return NEW_QUESTION

# def createChoiceQuestion()


def addQuestion(form_service, form, NEW_QUESTION):
    # Adds the question to the form
    question_setting = form_service.forms().batchUpdate(formId=form["formId"], body=NEW_QUESTION).execute()

    # Prints the result to show the question has been added
    get_result = form_service.forms().get(formId=form["formId"]).execute()
    print(get_result)