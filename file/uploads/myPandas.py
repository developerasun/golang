#!/usr/bin/env python
# coding: utf-8

# In[5]:


import pandas as pd 
# Pandas data structure: 1) Series(1 dimensional array), Dataframe() 


# In[6]:


obj = pd.Series([4,7,-5,3])


# In[7]:


obj


# In[8]:


# Extract values of the Series 
obj.values


# In[9]:


obj.keys


# In[10]:


obj[2]


# In[11]:


obj2 = pd.Series([1,2,3,4], index = ['a', 'x', 'o', 'd'])


# In[12]:


obj2['a']


# In[13]:


obj2['o'] = 99


# In[14]:


obj2


# In[15]:


obj2.isnull()


# In[16]:


data = {
        'city': ['Seoul', 'Busan', 'Daegu', 'Ulsan'], 
        'year': [2010, 2011, 2012, 2013]
       }


# In[17]:


frame = pd.DataFrame(data)


# In[18]:


frame


# In[19]:


sortedFrame = pd.DataFrame(data, columns=['year', 'city'])


# In[20]:


sortedFrame


# In[21]:


sortedFrame['year']


# In[22]:


sortedFrame['city']


# In[23]:


frame3 = pd.DataFrame(data, index = ['one', 'two', 'three', 'four'])


# In[24]:


frame3


# In[25]:


# Use func.loc instead of func.ix, ix is removed in Pandas ver 1.2.4
frame3.loc['two']


# In[26]:


pd.__version__


# In[27]:


transposedFrame3 = frame3.T


# In[28]:


transposedFrame3


# In[29]:


transposedFrame3['three']


# In[30]:


transposedFrame3.loc['city']


# In[31]:


transposedFrame3.loc['year']


# In[32]:


# axis=0 means along "indexes". It's a row-wise operation.
# axis=1 means along "columns". It's a column-wise operation.
transposedFrame3.drop(['two', 'four'], axis=1)


# In[33]:


transposedFrame3.drop(['year'], axis=0)


# In[40]:


frame4 = pd.DataFrame([[1,5,2,8], [3,4,1,8]], index = ['row2', 'row1'], columns = ['z', 'y', 'x', 'w'])


# In[41]:


frame4


# In[42]:


frame4.drop(['row1'], axis=0)


# In[43]:


frame4.drop(['row2'], axis=0)


# In[44]:


frame4.sort_index()


# In[45]:


frame4.sort_index(axis=1)


# In[52]:


frame4.sort_index(axis=1, ascending = False)


# In[56]:


frame4.sort_values(by=['y'])


# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:





# In[ ]:




