'use strict';

/**
 * @ngdoc function
 * @name dashboardApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the dashboardApp
 */
angular.module('dashboardApp')
  .controller('MainCtrl', ['$scope', '$websocket', '$interval', function ($scope, $websocket, $interval) {
    var timer, interval, updateTimer, secondsToHHMMSS;
    $scope.clock = "00:00:00";
    timer = 0;

    secondsToHHMMSS = function(value) {
      var seconds = parseInt(value % 60);
      var minutes = parseInt((value / 60) % 60);
      var hours = parseInt((value / (60 * 60)) % 24);
      hours = (hours < 10) ? "0" + hours : hours;
      minutes = (minutes < 10) ? "0" + minutes : minutes;
      seconds = (seconds < 10) ? "0" + seconds : seconds;
      return hours + ":" + minutes + ":" + seconds;
    };

    updateTimer = function() {
        timer += 1;
        $scope.clock = secondsToHHMMSS(timer);
    };

    $scope.teams = [];
    $scope.connected = false;
    $scope.running = false;
    $scope.solarFlare = false;

    $(".temperature").knob({
        'min': -142,
        'max': 35,
        'angleArc': 250,
        'angleOffset': -125,
        'readOnly': true,
        'fgColor': '#f56954',
    });

    $(".radiation").knob({
        'min': 0,
        'max': 1000,
        'angleArc': 250,
        'angleOffset': -125,
        'readOnly': true,
        'fgColor': '#f89406',
    });

    var ws = $websocket.$new({
      url: 'ws://localhost:8080/ws',
      protocols: []
    });

    ws.$on('$open', function () {
        console.log('Websocket is open');
        $scope.$apply(function() {
          $scope.connected = true;
        });
    });

    ws.$on('$message', function (data) {
        //console.log('Data received:', data);

        if ($scope.running !== data.running) {
          if (data.running) {
            timer = 0;
            interval = $interval(updateTimer, 1000);
          }
          else {
            $interval.cancel(interval);
          }
        }

        $scope.$apply(function() {
          $scope.running = data.running;
          $scope.solarFlare = data.readings.solarFlare;
        });

        $('.temperature')
          .val(data.readings.temperature)
          .trigger('change');

        $('.radiation')
          .val(data.readings.radiation)
          .trigger('change');

        if ( angular.toJson($scope.teams) !==  angular.toJson(data.teams) ) {
          $scope.$apply(function() {
            $scope.teams = data.teams;
          });
        }

    });

    ws.$on('$close', function () {
        console.log('Websocket closed!');
        $scope.$apply(function() {
          $scope.connected = false;
        });

    });
  }]);
