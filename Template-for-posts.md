### How to Git going
###### Getting started with git for wiki submissions (best practices)

This template should help you contribute to OMG Wiki. There's a shorter empty template here [http://omgcoporative.github.com]. It's fairly self explanatory, just download the primer and replace this text with a brief explanation on your tutorial. Followed by a list of key areas the article is focused on. It's always good to have a list of contents early on in the post to make sure that your readers have a good bearing. Contents can be either an ordered or unordered list.

Main content, relief and conclusions follow right after the contents list. After which you can post references. 

1. Note on structure
   - A short introduction (purpose)
   - Who the post is meant for
   - What you will learn 
   - Hang ups

2. Context (Submit from your local machine )

To get an overall idea of how git pull, merge and push works, read up here [http://]
For writing tips (tips on timely posts, relief, using pictures, media and lists) read here [http://]
	
	- For more help and discussion please join the OMGCo Slack group.

End of three part post

3. Resources
4. Conclusions and revision.

####### Note on structure
  
  - Introduction

  A short introduction goes here. The first few lines are important in getting the readers attention. Keep it simple and explain what the post is about. A single line or two should do for an introduction. For example, "This post is about posting your essays, tutorials and lessons on the Open Media Group Wiki. It's a guide for you to get git hooked up to your wiki branch effectively. Of course submitting to the wiki is easiest still on your browser."

  - Who this is meant for

  The Open Media Group wiki is a curriculum based on subjects for career programmers and newbies written by programmers. Posts are categorized by level desired. Its assumed that entry level programming tutorials and Computer Science posts will be most popular. But you can pretty much write about anything programming, tech, popular science or computer science related. There is a seperate post about writing specific curriculum posts here [http://omgcoporative.github.com] To access current posts you can either sign up for a github account and join the Society, or contribute as a guest. 
  
  This template is meant for contributions in general (Wiki, tutorials, curriculum, general tips...etc). We will cover git basics, how to contribute and how to post from your machine locally. Its assumed that you are somewhat familiar with git basics. You can also simplify this post by writing seperate guidelines and then linking back here. You can always talk to us about simplifying a specific post and find more about contributions by requesting an invite on Slack here [http://omgco.slack.com]


  - What you will learn
  If you are new to git, please head over to github.com and make an account. Its assumed that you have a working knowledge of Git. This post will cover the bare specifics of pointing a wiki branch on your local machine and then pushing new posts onto the OMG wiki.

  - Hang ups
  Your post may be merged with other edits, and any changes you make can also be merged to a previous commit. If you feel that your post loses its initial goals as decribed in the introduction, please make a pull request and reference back to the issue. Issues can be posted by anyone. Merging branches can have specific guides based on a particular curriculum. Please make sure that you go through Curriculum guides before posting. We keep a local copy of all posts as they first appear. Sometimes this is helpful to pull back a revision. In any case its best not to repeat yourself and create short, simple content (no more than 300 lines, or a 20 min read). **If your post gets too long, fire-up a seperate post and link back**.

####### Submit from your local machine

If you've worked with Sublime text and Git before, submitting and editing your wiki posts can help you maintain a workflow. Specially if you are continuing a series of tutorials or lessons. 

Assuming you are submitting to the main wiki, you can first create a directory and intialiaze Git. You can then pull the wiki/master branch.

	```

	$ mkdir mywikidir
	$ cd mywikidir
	$ git init

	$ git remote add origin mywikipost git@github.com:OMGCorporative/OMGwiki.wiki.git
	$ git remote
	mywikipost
	origin
	```

This effectively creates a remote called mywikipost on your machine and points it to our wiki HEAD:master branch. You can now either pull a file, folder or all wiki posts and make a local copy. Yes, we're giving it to you free, please use it responsibly (see creative commons license [http://www.omgcorporative.github.com/OMGwiki/wiki]).

Now that you're ready to submit your post, you can initialize a branch for that. First check if that's already been done for you by fetching the master branch.

	```
	$ git read-tree --prefix=wikipages/ -u wiki
	$ cd wikipages
	$ git fetch wiki
	$ git diff-tree -p wiki/master
	```

Check that you have a branch set up with 'git branch'.
If that doesn't work. Use 'git pull'.

You should now have a copy of posts on your machine. To make a new post you should open up a text editor and point it to that folder. In this case we're going to assume that you have Sublime Text set up. If you haven't done it already, read on [how to set up Sublime Text] before continuing.

When you are done editing or writing a new post. Commit your changes and push back to the wiki repo. Remember to check your status with 'git status', and commit any changes before you push.

	``` 
	$ git status
	$ git add .
	$ git commit -am "new changes"
	$ git push wiki HEAD:master
	```

Your changes should be reflected back on the wiki now. Your post is yours to write; add humor and ease into it. Make it as simple as you can. There's really no need to use big words unless your reader is already comfortable with your posts. Ratings for posts get people interested in what you have to say. We let our readers and students rate our content and suggest improvements. Remember, readers love it when you relate your own experiences.

####### Resources
These docs are great for understanding how Git subtrees work.
http://git-memo.readthedocs.org/en/latest/subtree.html

For an in-depth look at Git markdown read the docs at 
https://help.github.com/articles/basic-writing-and-formatting-syntax/

For tips on how to write a killer first paragraph read
http://essayheaven.blogspot.com/2009/11/writing-killer-first-paragraph.html

###### Conclusion



