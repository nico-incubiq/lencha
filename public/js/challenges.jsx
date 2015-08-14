
var Challenge = React.createClass({
    render: function() {
        return (
            <div className="col-sm-4">
                <div className="js-challenge challenge card">
                    <div className="content content-border">
                        <span className="js-title title">{this.props.challenge.title}</span>
                        <p className="js-decription">{this.props.challenge.smallDescription}</p>
                    </div>
                    <div className="action">
                        <a href={"/api/problems/" + this.props.challenge.apiUrl}>Solve</a>
                    </div>
                </div>
            </div>
        );
    }
});

var ChallengeList = React.createClass({
    getInitialState: function() {
        return {hiddenIds: []};
    },

    componentDidMount: function() {
        var ctx = this;
        var searchInput = document.getElementById('js-filter-challenges-input');
        searchInput.addEventListener('input', function() {
            ctx.onSearchInput(searchInput.value);
        });
    },

    render: function() {
        var ctx = this;
        var challenges = this.props.challenges.filter(function (challenge) {
            return ctx.state.hiddenIds.indexOf(challenge.id) == -1;
        });

        var rows = [];
        for(var i = 0; i < challenges.length; i+= 3) {
            var challengesComponents = [];
            for(var j = i; j < i + 3; j++) {
                if(j < challenges.length) {
                    challengesComponents.push(
                        <Challenge key={challenges[j].id} challenge={challenges[j]} />
                    );
                }
            }
            rows.push(
                <div className="row challenge-row">
                    {challengesComponents}
                </div>
            );
        }

        return (
            <div className="container">
                {rows}
            </div>
        );
    },

    searchInString: function(str, search) {
        return str.toLowerCase().search(search.toLowerCase()) == -1 ? false : true;
    },

    onSearchInput: function(searchValue) {
        var hiddenIds = this.props.challenges
                            .filter(function(challenge) {
                                return !this.searchInString(challenge.title, searchValue) && !this.searchInString(challenge.smallDescription, searchValue);
                            }, this)
                            .map(function(challenge) {
                                return challenge.id;
                            });

        this.setState({
            hiddenIds: hiddenIds
        });
    }
});

React.render(
  <ChallengeList challenges={window.challenges} />,
  document.getElementById('js-content')
);
