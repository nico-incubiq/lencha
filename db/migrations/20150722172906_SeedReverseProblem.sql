
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO problems (name, small_description, description, api_url)
 VALUES (
    'Reverse',
    'Reverse a string. Pretty simple isn''t it ? Solve this problem to activate your account.',
    '<h2>Description</h2>
<p>Ask the server for a string, reverse it and send it back, You’re done! One of the simplest problem, perfect to get you started.</p>
<h2>Solving Steps</h2>
<h4>1 - GET api/problems/reverse</h4>
<p>Gives you the problem data, in JSON. The string to reverse is in message.</p>
<pre><code>{
    "<span class="hljs-attribute">name</span>": <span class="hljs-value"><span class="hljs-string">"reverse"</span></span>,
    "<span class="hljs-attribute">status</span>": <span class="hljs-value"><span class="hljs-string">"in progress"</span></span>,
    "<span class="hljs-attribute">startedAt</span>": <span class="hljs-value"><span class="hljs-string">"2015-08-19T12:13:49.749778859+02:00"</span></span>,
    "<span class="hljs-attribute">endingAt</span>": <span class="hljs-value"><span class="hljs-string">"2015-08-19T12:14:19.749778925+02:00"</span></span>,
    "<span class="hljs-attribute">message</span>": <span class="hljs-value"><span class="hljs-string">"abcdefgh"</span>
</span>}
</code></pre>
<h4>2 - POST api/problems/reverse</h4>
<p>Once you have the reversed string, post this JSON.</p>
<pre><code>{
    "<span class="hljs-attribute">reversed</span>":<span class="hljs-value"><span class="hljs-string">"hgfedcba"</span>
</span>}
</code></pre>',
    'reverse'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM problems WHERE name='Reverse';
