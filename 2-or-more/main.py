from os import walk, listdir, path

import sys

# adyen
# affirm
# americommerce
# box
# checkbook
# datev
# enigma
# exactonline
# freeagent
# google_adsense
# google_adwords
# google_analytics
# holded
# liftoff
# mailchimp
# mixpanel
# nimble
# pabbly_emails
# pabbly_subscriptions
# paywhirl
# revcent
# storehippo
# xcart
# zoho
# zohobooks


def main():
    folders = sys.stdin
    # print(f"path {path}")

    # for dirname in next(walk(path))[1]:
    #     subfolder = path + "/" + dirname
    # dir_path = path.dirname(path.realpath(__file__))
    already_integrated = [
        "worldpay",
        "xero",
        "clearbooks",
        "braintree",
        "billomat",
        "yodlee",
        "authorizenet",
        "stripe",
        "klarna",
        "square",
        "shopify",
        "fastbill",
        "mx",
        "recurly",
        "quickbooks",
        "freshbooks",
        "enigma",
        "plaid",
    ]
    actually_oauth = [
        "patreon",
        "zoho",
        "zohobooks",
        "americommerce",
        "box",
        "checkbook",
        "exactonline",
        "freeagent",
        "storehippo",
    ]
    nope = ["task", "example"]
    dir_path = "/Users/adrianforsius/code/captec/connectors/stacks_connectors/integrations"
    for f in folders:
        f = f.rstrip()
        if f in already_integrated:
            continue

        if f in actually_oauth:
            continue

        if f in nope:
            continue

        contents = listdir(f"{dir_path}/{f}")
        if len(contents) > 3:
            print(f)


if __name__ == "__main__":
    main()
