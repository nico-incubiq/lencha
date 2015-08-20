
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO problems (name, small_description, description, api_url)
 VALUES (
    'Equation',
    'Solve a simple equation of the form : a*x^2 + b*x + c = 0',
    '<h2>Description</h2>
<p>Ask the server for a quadratic equation of the form: <span class="math inline"><span class="katex"><span class="katex-mathml"><math><semantics><mrow><mi>a</mi><msup><mi>x</mi><mn>2</mn></msup><mo>+</mo><mi>b</mi><mi>x</mi><mo>+</mo><mi>c</mi><mo>=</mo><mn>0</mn></mrow><annotation encoding="application/x-tex">ax^2 + bx + c = 0</annotation></semantics></math></span><span class="katex-html" aria-hidden="true"><span class="strut" style="height:0.8141079999999999em;"></span><span class="strut bottom" style="height:0.897438em;vertical-align:-0.08333em;"></span><span class="base textstyle uncramped"><span class="mord mathit">a</span><span class="mord"><span class="mord mathit">x</span><span class="vlist"><span style="top:-0.363em;margin-right:0.05em;"><span class="fontsize-ensurer reset-size5 size5"><span style="font-size:0em;">​</span></span><span class="reset-textstyle scriptstyle uncramped"><span class="mord">2</span></span></span><span class="baseline-fix"><span class="fontsize-ensurer reset-size5 size5"><span style="font-size:0em;">​</span></span>​</span></span></span><span class="mbin">+</span><span class="mord mathit">b</span><span class="mord mathit">x</span><span class="mbin">+</span><span class="mord mathit">c</span><span class="mrel">=</span><span class="mord">0</span></span></span></span></span>.
Solve it, and send it back. We crafted the equation so the solutions are two integers. Once you have
the answer, send it back to the server. You’re done!</p>
<h2>Solving Steps</h2>
<h4>1 - GET api/problems/equation</h4>
<p>Gives you the problem data, in JSON. The string representing the equation is the message attribute.</p>
<pre><code>{
    "<span class="hljs-attribute">name</span>": <span class="hljs-value"><span class="hljs-string">"equation"</span></span>,
    "<span class="hljs-attribute">status</span>": <span class="hljs-value"><span class="hljs-string">"in progress"</span></span>,
    "<span class="hljs-attribute">startedAt</span>": <span class="hljs-value"><span class="hljs-string">"2015-08-19T12:13:49.749778859+02:00"</span></span>,
    "<span class="hljs-attribute">endingAt</span>": <span class="hljs-value"><span class="hljs-string">"2015-08-19T12:14:19.749778925+02:00"</span></span>,
    "<span class="hljs-attribute">message</span>": <span class="hljs-value"><span class="hljs-string">"847*x^2+27104*x-690434591=0"</span>
</span>}
</code></pre>
<h4>2 - POST api/problems/equation</h4>
<p>Once you solved the polynomial equation, post this JSON.</p>
<pre><code>{
    "<span class="hljs-attribute">x1</span>":<span class="hljs-value"><span class="hljs-string">-919</span>
    "<span class="hljs-attribute">x2</span>":<span class="hljs-value"><span class="hljs-string">887</span>
</span>}
</code></pre>',
    'equation'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM problems WHERE name='Equation';
