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

    var data2 = {
        labels: ["January", "February", "March", "April", "May", "June", "July"],
        datasets: [
            {
                label: "My First dataset",
                fillColor: "rgba(220,220,220,0.2)",
                strokeColor: "rgba(220,220,220,1)",
                pointColor: "rgba(220,220,220,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(220,220,220,1)",
                data: [65, 59, 80, 81, 56, 55, 40]
            },
            {
                label: "My Second dataset",
                fillColor: "rgba(151,187,205,0.2)",
                strokeColor: "rgba(151,187,205,1)",
                pointColor: "rgba(151,187,205,1)",
                pointStrokeColor: "#fff",
                pointHighlightFill: "#fff",
                pointHighlightStroke: "rgba(151,187,205,1)",
                data: [28, 48, 40, 19, 86, 27, 90]
            }
        ]
    };

    var ctx = document.getElementById("myChart").getContext("2d");
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

        myNewChart.addData([data.radiation, 10], "August");
        myNewChart.removeData();
        console.log(data2.datasets[0].data);
        //myNewChart.update();
    });

    ws.$on('$close', function () {
        console.log('Websocket closed!');
    });
  });
