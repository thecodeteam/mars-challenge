#!/usr/bin/env python

import requests  # https://github.com/kennethreitz/requests/
import time
import random

server_url = 'http://192.168.59.103:8080/api'
admin_header = {'X-Auth-Token': '1234'}


# Server Calls ----------------------------------------------
def server_start(server_url):
    """Start game
       curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
    """
    url = server_url + '/start'
    start_game = requests.post(url, headers=admin_header)
    if start_game.status_code == 200:
        print('Server: Game has been Started!')
    else:
        print ('Server: Game Start Failed!')
        print ("HTTP Code: " + str(start_game.status_code) + " | Response: " + start_game.text)


def server_reset(server_url):
    """Reset game
       curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
    """
    url = server_url + '/reset'
    reset_game = requests.post(url, headers=admin_header)
    if reset_game.status_code == 200:
        print('Server: Game has been Reset!')
    else:
        print ('Server: Game Reset Failed!')
        print ("HTTP Code: " + str(reset_game.status_code) + " | Response: " + reset_game.text)


def server_stop(server_url):
    """Stop game
       curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/stop
    """
    url = server_url + '/stop'
    stop_game = requests.post(url, headers=admin_header)
    if stop_game.status_code == 200:
        print('Server: Game has been Stopped!')
    else:
        print ('Server: Game Stop Failed!')
        print ("HTTP Code: " + str(stop_game.status_code) + " | Response: " + stop_game.text)


def server_check_game_started(server_url):
    """Start game
       curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
    """
    url = server_url + '/start'
    start_game = requests.post(url, headers=admin_header)
    if start_game.status_code == 400:
        return True
    else:
        return False

# Shield Calls ------------------------------------------------

#Shield Manipulation
def team_shield_up(team_name, team_auth):
    """
    Sets the team shield up
    curl -i -H 'X-Auth-Token: 1335aa6af5d0289f' -X POST http://localhost:8080/api/shield/enable

    """
    url = server_url + '/shield/enable'
    auth_header = {'X-Auth-Token': team_auth}
    shield_up = requests.post(url, headers=auth_header)
    if shield_up.status_code == 200:
        print ('Server: Team: ' + team_name + ' Shield is UP!')
    else:
        print ('Server: Team: ' + team_name + ' Shield UP! request Failed!')
        print ("HTTP Code: " + str(shield_up.status_code) + " | Response: " + shield_up.text)

#Starting Shield Manipulation
def team_shield_down(team_name, team_auth):
    """
    Sets the team shield Down
    curl -i -H 'X-Auth-Token: 1335aa6af5d0289f' -X POST http://localhost:8080/api/shield/disable
    """
    url = server_url + '/shield/disable'
    auth_header = {'X-Auth-Token': team_auth}
    shield_down = requests.post(url, headers=auth_header)
    if shield_down.status_code == 200:
        print ('Server: Team: ' + team_name + ' Shield is DOWN!')
    else:
        print ('Server: Team: ' + team_name + ' Shield DOWN! request Failed!')
        print ("HTTP Code: " + str(shield_down.status_code) + " | Response: " + shield_down.text)


print("Starting the Test Harness")
print("-------------------------")

print("\nChecking the Server Status...")
# Check that Server has not started
if server_check_game_started(server_url):
    print("...previous game running....")
    server_stop(server_url)
    server_reset(server_url)
else:
    print("...cleaning up....")
    server_stop(server_url)
    server_reset(server_url)

time.sleep(2)

print("\nAdding Teams...")
# Adding team: TheBorg
team1_name = 'TheBorgs'
team1_auth = ''

url = server_url + '/join/' + team1_name
payload = ''
# POST with form-encoded data
response = requests.post(url, data=payload)

team1_auth = response.text

if response.status_code == 200:
    print ('Team \'' + team1_name + '\' joined the game!')
    print (team1_name + ' authentication Code: ' + team1_auth)
else:
    print ('Team \'' + team1_name + '\' joining game Failed!')
    print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)
time.sleep(2)

# Adding team: QuickFandango
team2_name = 'QuickFandango'
team2_auth = ''

url = server_url + "/join/" + team2_name
payload = ''
# POST with form-encoded data
response = requests.post(url, data=payload)

team2_auth = response.text

if response.status_code == 200:
    print ('Team \'' + team2_name + '\' joined the game!')
    print (team2_name + ' authentication Code: ' + team2_auth)
else:
    print ('Team \'' + team2_name + '\' joining game Failed!')
    print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)
time.sleep(2)

# Adding team: InTheBigMessos
team3_name = 'InTheBigMessos'
team3_auth = ''

url = server_url + "/join/" + team3_name
payload = ''
# POST with form-encoded data
response = requests.post(url, data=payload)

team3_auth = response.text

if response.status_code == 200:
    print ('Team \'' + team3_name + '\' joined the game!')
    print (team3_name + ' authentication Code: ' + team3_auth)
else:
    print ('Team \'' + team3_name + '\' joining game Failed!')
    print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)

time.sleep(10)
print("\nStarting the Game...")
# Starting the Game Server
server_start(server_url)

# Starting the Teams Logic
while True:
    team = random.randrange(0,3)
    action = random.randrange(1,3)
    team_list =[(team1_name,team1_auth), (team2_name, team2_auth), (team3_name, team3_auth)]
    # print("\nGameMove: Team: " + team_list[team][0] + ' Action:' + str(action) + ' Name: ' + team_list[team][0] +'|'+ team_list[team][1])
    if action > 1:
        print("\nGameMove: Team: " + team_list[team][0] + ' Action: Shield UP! | Team Key: ' + team_list[team][1])
        team_shield_up(team_list[team][0],(team_list[team][1]))
    else:
         print("\nGameMove: Team: " + team_list[team][0] + ' Action: Shield Down! | Team Key: ' + team_list[team][1])
         team_shield_down(team_list[team][0],(team_list[team][1]))
    time.sleep(2)
    if server_check_game_started(server_url) == False:
        print('\nServer: Game is Over...')
        break















