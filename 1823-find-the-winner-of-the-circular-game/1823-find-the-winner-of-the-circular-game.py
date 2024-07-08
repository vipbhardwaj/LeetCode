class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        players = list(range(1, n+1))
        i = 1
        while len(players) > 1:
            i += k-1
            while i > len(players):
                i -= len(players)
            # print(f"now i = {i}")
            players.pop(i-1)
            # print(players)
        return players[0]