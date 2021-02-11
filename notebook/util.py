def abbreviate(s: str) -> str:
    words = s.split('-')
    return ''.join(map(lambda x: x[0].upper(), words))
