
(function() {
    var jumbotron = d3.select('#js-jumbotron');
    var svg = d3.select('#js-jumbotron-svg-background');

    // Setting the correct height of the svg
    var jumbotronHeight = jumbotron.node().getBoundingClientRect().height;
    svg.style('height', jumbotronHeight + 'px');

    var width  = svg.node().getBoundingClientRect().width;
    var height = svg.node().getBoundingClientRect().height;

    console.log('Svg width: ' + width + ' height: '+ height);

    // Sort algorithm
    var data = _.shuffle(_.range(100));

    var bars = svg.selectAll('.sort-bar');
    var barsUpdate = bars.data(data);
    var barsEnter  = barsUpdate.enter();

    barsEnter.append('svg:line')
        .attr('class', 'sort-bar')
        .attr('x1', function(d, i) { return width * (data.length - i) / data.length; })
        .attr('y1', function(d) { return height; })
        .attr('x2', function(d, i) { return width * (data.length - i) / data.length; })
        .attr('y2', function(d) { return height - height * (data.length - d) / data.length; });
})();
