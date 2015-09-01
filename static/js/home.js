
// JUMBOTRON BACK GROUND
(function() {

    var delayTransition = 3 * 1000;

    function genConstData(length) {
        var data = [];
        for(var x = 0; x < length; x++) {
            data[x] = 0.5;
        }
        return data;
    }

    function genWaveData(dec) {
        return function(length) {
            var data = [];
            for(var x = 0; x < length; x++) {
                data[x] = (Math.sin((x / length) * 2 * Math.PI + dec) + 1) / 2;
            }
            return data;
        };
    }

    function updateJumbotronBackground(data, cb) {
        var bars = svg.selectAll('.sort-bar');
        var barsUpdate = bars.data(data);
        var barsEnter  = barsUpdate.enter();
        var barsExit   = barsUpdate.exit();

        barsUpdate.transition().duration(delayTransition)
            .attr('class', 'sort-bar update')
            .attr('data', function(d) { return d; })
            .attr('y1', function(d) { return height; })
            .attr('y2', function(d) { return height - height * d; });

        barsEnter.append('svg:line')
            .attr('class', 'sort-bar')
            .attr('data', function(d) { return d; })
            .attr('x1', function(d, i) { return (i / data.length) * width; })
            .attr('x2', function(d, i) { return (i / data.length) * width; })
            .attr('y1', function(d) { return height; })
            .attr('y2', function(d) { return height - height * d; });
    }

    var jumbotron = d3.select('#js-jumbotron');
    var svg = d3.select('#js-jumbotron-svg-background');

    // Setting the correct height of the svg
    var jumbotronHeight = jumbotron.node().getBoundingClientRect().height;
    svg.style('height', jumbotronHeight + 'px');

    var width  = svg.node().getBoundingClientRect().width;
    var height = svg.node().getBoundingClientRect().height;

    console.log('svg width: ' + width + ' height: '+ height);

    var numberOfBars = Math.floor(width / 5.0);

    var generators = [genWaveData(0), genWaveData(Math.PI / 2), genWaveData(Math.PI)];

    var curGen = 0;
    function loop() {
        data = generators[curGen](numberOfBars);
        updateJumbotronBackground(data);
        curGen = (curGen + 1) % generators.length;

        setTimeout(function() {
            loop();
        }, delayTransition + 500);
    }

    updateJumbotronBackground(genWaveData(0)(numberOfBars));
    curGen = (curGen + 1) % generators.length;
    setTimeout(loop, 500);
})();

(function() {
    var examples = document.getElementsByClassName('code-example');
    var exampleLinks = document.getElementById('js-language-links').getElementsByTagName('a');

    // Handle clicks
    _.forEach(exampleLinks, function(link, i){
        link.addEventListener('click', function(evt){
            evt.preventDefault();

            _.forEach(exampleLinks, function(link, j) {
                link.className = i == j ? 'selected' : '';
            });

            _.forEach(examples, function(example, j) {
                example.style.display = i == j ? 'block' : 'none';
            });
        });
    });

    // Active first link
    examples[0].style.display = 'block';
    exampleLinks[0].className = 'selected';
})();

