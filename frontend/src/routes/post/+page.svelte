<script lang="ts">
	let right_img:String = "right.png";
	let wrong_img:String = "wrong.png";
</script>

<h1>My thoughts on UML data diagrams</h1>

<p>I have recently started using UML data diagrams again for a project for university.</p>
<p>It's been a while since I last used UML for any project. I never use it for my personal projects, because I always thought that this was a waste of time. Because how on earth is someone going to define the whole structure of a serviced before they actually start creating the service.</p>

<p>From the (limited) experience with UML, I know that I always end up discarding the whole diagram, because in practice this beautifully thought out design does not work. As in, it is technically not possible to implement the diagram in code.</p>

<h2>UML provides clarity</h2>

<p>However, recent event have led me to use UML diagrams again. Not only in university assignments, but also in personal projects that were too complex to be designed in my head and a needed clear, structured overview that I could implement in code.</p>

<p>When designing these diagrams I could see problems arising before I started coding the service.</p>
<p>What often happend in the past is, I would not foresee any problems untill I actually began to implement the components of the service. And when I saw a problem arise, I would often need to throw away hours of work and redo/rethink that part of the system.</p>
<p>So in these cases my UML diagrams, which I used to despise, now save me a lot of time and more importantly a lot of frustation.</p>

<h2>UML improves communication</h2>

<p>I have also seen an application for UML besides designing projects.</p>
<p>In a couple projects it has helped me clarify my thought to other people in my team.</p>
<p>In a way, these diagrams allow you to clearly communicate the design that you have in your head.</p>
<p>Telling someone what they should make is not always very easy. People can misunderstand you when you explain your design using words, which can lead to people disagree unnecessarily with you or implement something different from what you had in mind.</p>

<h2>My UML design practices</h2>

<p>UML diagrams are still, in my opinion, not a correct design of the eventual system.</p>
<p>But I now think they can provide my with a lot of guidance when implementing a system.</p>
<p>My recent UML diagrams do not have an exact design of al the classes that will be present in the code.</p>
<p>I have moved away from definig all public and private properties for every class. This was the way that I was thought.</p>
<p>But I found it much more helpful to only define the necessary properties in a class that the user of the system needs. So not properties that are only necessary for the system to work.</p>
<p>Like userId, which is useful for a system to be able to reference to a user object.</p>
<p>But a user is probably not interested in this property, but is instead interested in something like a username.</p>
<p>So I now only put properties in my diagrams that are needed for the user of the service.</p>
<p>The systemproperties are going to be different anyway when I start coding. Because who knows, maybe I don't even need this userId property and then all that effort I spent on that part of the diagram is wasted.</p>
<p>Creating UML data diagrams in this way let me to implement something without breaking the design and wasting too much time.</p>

<p>Below is an example of a diagram where on the left I used classes with all private and public properties and on the right I used a simplified version of the diagram.</p>
<p>In the simplified version some properties are not written down, but can be derived from the arrows in the diagram.</p>

<div class="img_row">
	<img src="wrong.png" alt="UML data diagram with all the properties"/>
	<img src="right.png" alt="UML data diagram in simplified format with only the propeties that the user cares about"/>
</div>

<style>
	* {
		font-family: sans-serif;
	}
</style>
