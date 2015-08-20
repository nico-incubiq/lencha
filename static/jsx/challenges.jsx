(function() {

    var Challenge = React.createClass({

        render: function() {
            var solved;
            if(this.props.challenge.solved) {
                solved = <span className="solved pull-right text-success ion-checkmark"> Solved</span>
            }

            return (
                <div className="col-sm-4">
                    <div className="js-challenge challenge card">
                        <div className="content content-border">
                            <span className="title">{this.props.challenge.id}. {this.props.challenge.name} {solved}</span>
                            <p>{this.props.challenge.smallDescription}</p>
                        </div>
                        <div className="action">
                            <span className="pull-left">Solved By: {this.props.challenge.solvedTotal}</span>
                            <a href={"/problems/" + this.props.challenge.apiUrl}>Solve</a>
                        </div>
                    </div>
                </div>
            );
        }
    });

    var ChallengeList = React.createClass({
        getInitialState: function() {
            return {
                showSolved: true,
                hiddenIds: []
            };
        },

        componentDidMount: function() {
            var ctx = this;
            var searchInput = document.getElementById('js-filter-challenges-input');
            searchInput.addEventListener('input', function() {
                ctx.onSearchInputChange(searchInput.value);
            });

            var solvedCheckbox = document.getElementById('js-filter-solved-checkbox');
            solvedCheckbox.addEventListener('click', function() {
                ctx.onSolvedCheckboxClick(solvedCheckbox.checked);
            });
        },

        render: function() {
            var ctx = this;
            var challenges = this.props.challenges.filter(function (challenge) {
                return ctx.state.hiddenIds.indexOf(challenge.id) == -1;
            });

            if(!this.state.showSolved) {
                challenges = challenges.filter(function(challenge) {
                    return challenge.solved == false;
                });
            }

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

        onSearchInputChange: function(searchValue) {
            var hiddenIds = this.props.challenges
                                .filter(function(challenge) {
                                    return !this.searchInString(challenge.name, searchValue) && !this.searchInString(challenge.smallDescription, searchValue);
                                }, this)
                                .map(function(challenge) {
                                    return challenge.id;
                                });

            this.setState({
                hiddenIds: hiddenIds
            });
        },

        onSolvedCheckboxClick: function(checked) {
            this.setState({showSolved: checked});
        }
    });

    React.render(
      <ChallengeList challenges={window.challenges} />,
      document.getElementById('js-content')
    );
})();
