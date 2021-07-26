google.charts.load('current', {'packages':['corechart']});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {

  var data = google.visualization.arrayToDataTable([
    ['Name', 'How many marbles a person owns'],
    ['Taewoong',     5000],
    ['Hyesu',      8000],
    ['Junseok',  1250],
    ['Jonghyun', 700],

  ]);

  var options = {
    title: 'Kosk Marbles Ledger'
  };

  var chart = new google.visualization.PieChart(document.getElementById('piechart'));

  chart.draw(data, options);
}