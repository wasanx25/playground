describe('blogを開く', () => {
  before(() => {
    cy.visit('https://blog.wasanx25.com/')
  })

  it('タイトルが正常', () => {
    cy.get('#header > h1').contains('wasanx25 blog').should('not.be.hidden')
  })
})