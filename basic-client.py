#!/usr/bin/env python

import requests

from argparse import ArgumentParser

def main(doorman, indicator):

    r = requests.post(doorman,json=indicator)
    

if __name__ == '__main__':

    p = ArgumentParser('Submit a new indicator to doorman.')
    p.add_argument('doorman',
                    description='Service instance for doorman.')
    p.add_argument('indicator',
                    description='The indicator that you would like to add.')
    p.add_argument('reason',
                    description='Why this indicator is being added.')
    p.add_argument('-t', '--ttl',
                    default='1h',
                    description='How long the indicator should live for.')
    args = p.parse_args()

    indicator=dict(
        indicator=args.indicator,
        reason=args.reason,
        ttl=args.ttl
    )

    main(args.doorman, indicator)