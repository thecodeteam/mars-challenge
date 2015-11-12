#!/usr/bin/env python
# Test Harness to test the Game Controller Functionality
#


import requests  # https://github.com/kennethreitz/requests/
import time
import random
import os


# Global Variables
controller_endpoint = os.environ['GC_ENDPOINT']

#Replace this variables as appropriate
server_url = 'http://' + controller_endpoint + '/api'
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
    game_stop_url = server_url + '/stop'
    stop_game = requests.post(game_stop_url, headers=admin_header)
    if stop_game.status_code == 200:
        print('Server: Game has been Stopped!')
    else:
        print ('Server: Game Stop Failed!')
        print ("HTTP Code: " + str(stop_game.status_code) + " | Response: " + stop_game.text)


def server_check_game_started(server_url):
    """
    Start game
    curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
    """
    gstart_url = '{0}/start'.format(server_url)
    start_game = requests.post(gstart_url, headers=admin_header)
    if start_game.status_code == 400:
        return True
    else:
        return False


def server_kick_team(server_url, team_name):
    """
    Kicks a team from the registration list. Before the game is started.
    Example Curl: curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/kick/foobar
    :param server_url: The Server URL
    :param team_name: The Team's name
    """
    kick_url = server_url + '/kick/' + team_name
    team_kicked = requests.post(kick_url, headers=admin_header)
    if team_kicked.status_code == 200:
        print('Server: The team: {0} has been Kicked out!'.format(team_name))
    else:
        print ('Server: Team Kick failed for Team: {0}'.format(team_name))
        print ("HTTP Code: {0} | Response: {1}".format(str(team_kicked.status_code), team_kicked.text))


def server_config(server_url):
    """
    Retries the Server's configuration parameters
    curl -i -X GET http://localhost:8080/api/config
    :param server_url:
    :return: Nothing
    """
    kick_url = '{0}/config'.format(server_url)
    srv_config = requests.get(kick_url)
    print ("HTTP Code: {0} | Response: {1}".format(str(srv_config.status_code), srv_config.text))


# Shield Calls ------------------------------------------------
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


while True:

    # Test Harness ------------------------------------------------
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

    # Testing Server Configuration Functionality
    # ------------------------------------------------

    print("\nChecking the Server Configuration")
    print("------------------------------------")
    server_config(server_url)

    # Testing Adding Teams to Game Functionality
    # ------------------------------------------------

    print("\nAdding Teams")
    print("--------------")

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



    # Testing Kick Team Functionality
    # ------------------------------------------------

    print("\nChecking the Server Kick Functionality")
    print("----------------------------------------")
    print("Kicking {0}  team out...".format(team1_name))
    print("Team {0} has Auth Key: {1}".format(team1_name, str(team1_auth)))

    server_kick_team(server_url, team1_name)

    print("Adding  {0}  team back in...".format(team1_name))

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

    time.sleep(10)


    # Starting the the GAME
    # ------------------------------------------------

    print("\nStarting the Game")
    print("-------------------")
    # Starting the Game Server
    server_start(server_url)

    # Starting the Teams Logic
    while True:
        team = random.randrange(0, 1)
        action = random.randrange(1, 3)
        team_list = [(team1_name, team1_auth)]
        # print("\nGameMove: Team: " + team_list[team][0] + ' Action:' + str(action) + ' Name: ' + team_list[team][0] +'|'+ team_list[team][1])
        if action > 1:
            print("\nGameMove: Team: " + team_list[team][0] + ' Action: Shield UP! | Team Key: ' + team_list[team][1])
            team_shield_up(team_list[team][0], (team_list[team][1]))
        else:
            print("\nGameMove: Team: " + team_list[team][0] + ' Action: Shield Down! | Team Key: ' + team_list[team][1])
            team_shield_down(team_list[team][0], (team_list[team][1]))
        time.sleep(2)
        if server_check_game_started(server_url) == False:
            print('\nServer: Game is Over...')
            break
