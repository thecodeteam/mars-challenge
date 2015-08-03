'use strict';

/**
 * @ngdoc function
 * @name dashboardApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the dashboardApp
 */
angular.module('dashboardApp')
  .controller('MainCtrl', ['$scope', '$websocket', function ($scope, $websocket) {
    $scope.teams = [];
    $scope.connected = false;
    $scope.running = false;

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
        console.log('Data received:', data);

        $scope.running = data.running;
        $scope.solarFlare = data.readings.solarFlare;

        $('.temperature')
          .val(data.readings.temperature)
          .trigger('change');

        $('.radiation')
          .val(data.readings.radiation)
          .trigger('change');

        $scope.$apply(function() {
          $scope.teams = data.teams;
        });
    });

    ws.$on('$close', function () {
        console.log('Websocket closed!');
        $scope.$apply(function() {
          $scope.connected = false;
        });

    });
  }]);
