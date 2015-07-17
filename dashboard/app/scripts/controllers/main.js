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
    var maxHistory = 20;

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

    var data2 = {
        labels: [],
        datasets: [
            {
                label: "Temperature",
                fillColor: "rgba(220,220,220,0.2)",
                strokeColor: "rgba(220,220,220,1)",
                pointColor: "rgba(220,220,220,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(220,220,220,1)",
                data: []
            },
            {
                label: "Radiation",
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: []
            }
        ]
    };

    var ctx = document.getElementById("history").getContext("2d");
    var myNewChart = new Chart(ctx).Line(data2, {
        bezierCurve: false
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

        myNewChart.addData([data.temperature, data.radiation], "");
        if (myNewChart.datasets[0].points.length > maxHistory){
            myNewChart.removeData();
        }
    });

    ws.$on('$close', function () {
        console.log('Websocket closed!');
    });
  });
