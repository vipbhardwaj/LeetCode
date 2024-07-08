class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        players = list(range(1, n+1))
        i = 0
        while len(players) > 1:
            i = (i + k-1) % len(players)
            # i %= len(players)
            # print(f"now i = {i}")
            players.pop(i)
            # print(players)
        return players[0]