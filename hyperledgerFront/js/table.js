google.charts.load('current', {'packages':['table']});
google.charts.setOnLoadCallback(drawTable);

function drawTable() {
  var data = new google.visualization.DataTable();
  data.addColumn('string', 'Name');
  data.addColumn('number', 'Marbles($)');
  data.addColumn('boolean', 'Demo done?');
  data.addRows([
    ['Taewoong',  {v: 5000, f: '$5,000'}, true],
    ['Hyesu',   {v:8000,   f: '$8,000'},  true],
    ['Junseok', {v: 1250, f: '$1,250'}, true],
    ['Jonghyun',   {v: 700,  f: '$700'},  true]
  ]);

  var table = new google.visualization.Table(document.getElementById('table_div'));

  table.draw(data, {showRowNumber: true, width: '100%', height: '100%'});
}