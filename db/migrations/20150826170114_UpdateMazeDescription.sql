
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE problems
SET description='<h2><a></a>Description</h2>
<p>Solve a maze. The maze structure is as such :</p>
<table class="maze-table">
<tbody>
<tr>
<td>S</td>
<td>X</td>
<td>E</td>
<td>-</td>
<td>-</td>
</tr>
<tr>
<td>-</td>
<td>-</td>
<td>X</td>
<td>X</td>
<td>-</td>
</tr>
<tr>
<td>X</td>
<td>-</td>
<td>-</td>
<td>X</td>
<td>-</td>
</tr>
<tr>
<td>-</td>
<td>X</td>
<td>-</td>
<td>X</td>
<td>-</td>
</tr>
<tr>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td>-</td>
</tr>
</tbody>
</table>
<p>The X represent the walls, and the - an empty case. You start at the case which contains a S, and your goal is to end to the case containing a E. Below is the solution (each maze accept a unique solution).</p>
<table class="maze-table">

<tbody>
<tr>
<td>S</td>
<td>X</td>
<td>E</td>
<td>13</td>
<td>12</td>
</tr>
<tr>
<td>1</td>
<td>2</td>
<td>X</td>
<td>X</td>
<td>11</td>
</tr>
<tr>
<td>X</td>
<td>3</td>
<td>4</td>
<td>X</td>
<td>10</td>
</tr>
<tr>
<td>-</td>
<td>X</td>
<td>5</td>
<td>X</td>
<td>9</td>
</tr>
<tr>
<td>-</td>
<td>-</td>
<td>6</td>
<td>7</td>
<td>8</td>
</tr>
</tbody>
</table>
<p>To finish the problem, you’ll have to send the answer in the form of a string,
which can only contain 4 characters : U (up) D (down) R (right) L (left).</p>
<p>Here you would send : DRDRDDRRUUUULL</p>
<h2><a></a>Api</h2>
<h4><a></a>GET api/problems/maze</h4>
<p>Gives you the problem data, in JSON. The maze to solve is in message.</p>
<pre><code class="language-json">{
    "<span class="hljs-attribute">name</span>": <span class="hljs-value"><span class="hljs-string">"maze"</span></span>,
    "<span class="hljs-attribute">status</span>": <span class="hljs-value"><span class="hljs-string">"in progress"</span></span>,
    "<span class="hljs-attribute">startedAt</span>": <span class="hljs-value"><span class="hljs-string">"2015-08-26T10:34:46.089520466-04:00"</span></span>,
    "<span class="hljs-attribute">endingAt</span>": <span class="hljs-value"><span class="hljs-string">"2015-08-26T10:34:51.089520671-04:00"</span></span>,
    "<span class="hljs-attribute">message</span>": <span class="hljs-value">{
        "<span class="hljs-attribute">height</span>": <span class="hljs-value"><span class="hljs-number">5</span></span>,
        "<span class="hljs-attribute">width</span>": <span class="hljs-value"><span class="hljs-number">5</span></span>,
        "<span class="hljs-attribute">maze</span>": <span class="hljs-value">[
            <span class="hljs-string">"SXE--"</span>,
            <span class="hljs-string">"--XX-"</span>,
            <span class="hljs-string">"X--X-"</span>,
            <span class="hljs-string">"-X-X-"</span>,
            <span class="hljs-string">"-----"</span>
        ]
    </span>}
</span>}
</code></pre>

<h4><a></a>POST api/problems/maze</h4>
<p>The JSON you have to send</p>
<pre><code>{
    "<span class="hljs-attribute">solution</span>": <span class="hljs-value"><span class="hljs-string">"DRDRDDRRUUUULL"</span>
</span>}
</code></pre>
'
WHERE api_url='maze';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
UPDATE problems SET description='Find the only solution of a maze' WHERE api_url='maze';

