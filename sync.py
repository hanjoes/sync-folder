import argparse
import json


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='CLI controls how sync behaves')
    parser.add_argument('--src', required=True,
                        help=('source folder if bi is set'))
    parser.add_argument('--dst', required=True,
                        help=('destination folder if bi is set'))
    parser.add_argument('--bi', required=True,
                        help=('whether the synchronization is bi-directional'))

    args = parser.parse_args()
    print(args)
