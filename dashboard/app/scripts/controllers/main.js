'use strict';

/**
 * @ngdoc function
 * @name dashboardApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the dashboardApp
 */
angular.module('dashboardApp')
  .controller('MainCtrl', function ($websocket) {
    this.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];

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
        //'fgColor': '#f56954',
    });

    var ws = $websocket.$new({
      url: 'ws://localhost:8080/ws',
      protocols: []
    });

    ws.$on('$open', function () {
        console.log('Websocket is open');
    });

    ws.$on('$message', function (data) {
        console.log('Data received:', data);
        $('.temperature')
          .val(data.temperature)
          .trigger('change');

        $('.radiation')
          .val(data.radiation)
          .trigger('change');
    });

    ws.$on('$close', function () {
        console.log('Websocket closed!');
    });
  });
