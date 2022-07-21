# baseball-luck-twitbot
Application that calculates how lucky an MLB player has been for the current season and then tweets that information out. Written in Go.

Calculation based off of the magnitude of difference between the expected result of a player's plate appearances and the player's actual
results from plate appearances.

## pkg/stats ##

FindStat supports getting any of the following stats from a hitter :

G PA HR R RBI SB BB% K% ISO BABIP AVG OBP SLG wOBA xwOBA wRC+ BsR Off Def WAR
(source: fangraphs.com)

## main.go ##

This handles the deployment of the application and tweets out the result to https://twitter.com/luckmlb 


