google.charts.load('current', {packages:["orgchart"]});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {
  var data = new google.visualization.DataTable();
  data.addColumn('string', 'Name');
  data.addColumn('string', 'Manager');
  data.addColumn('string', 'ToolTip');

  // For each orgchart box, provide the name, manager, and tooltip to show.
  data.addRows([
    [{'v':'goBlock', 'f':'goBlock<div style="color:red; font-style:italic">Blockchain 6-month-long program</div>'},
     '', 'The Program'],
    [{'v':'Class A', 'f':'Class A<div style="color:red; font-style:italic">Best Class</div>'},
     'goBlock', 'class'],
    [{'v':'Team Kosk', 'f':'Team Kosk<div style="color:red; font-style:italic">Team 1</div>'},
     'Class A', 'team kosk'],
    ['Team Kosk', 'Class A', ''],
    ['Taewoong', 'Team Kosk', ''],
    ['Hyesu', 'Team Kosk', ''],
    ['Junseok', 'Team Kosk', ''],
    ['Jonghyun', 'Team Kosk', '']
 
  ]);

  // Create the chart.
  var chart = new google.visualization.OrgChart(document.getElementById('chart_div'));
  // Draw the chart, setting the allowHtml option to true for the tooltips.
  chart.draw(data, {'allowHtml':true});
}