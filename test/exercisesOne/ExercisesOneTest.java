package exercisesOne;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

class ExercisesOneTest {

    @Test
    void testTwoSum() {
        ExercisesOne exercisesOne = new ExercisesOne();

        int[] numsOne = {2, 7, 11, 15};
        int targetOne = 9;
        int[] expectedOne = {0, 1};
        assertArrayEquals(expectedOne, exercisesOne.twoSum(numsOne, targetOne));

        int[] numsTwo = {3, 2, 4};
        int targetTwo = 6;
        int[] expectedTwo = {1, 2};
        assertArrayEquals(expectedTwo, exercisesOne.twoSum(numsTwo, targetTwo));

        int[] numsThree = {0};
        int targetThree = 10;
        int[] expectedThree = {};
        assertArrayEquals(expectedThree, exercisesOne.twoSum(numsThree, targetThree));

        int[] numsFour = {-1, -2, -3, -4, -5};
        int targetFour = -8;
        int[] expectedFour = {2, 4};
        assertArrayEquals(expectedFour, ExercisesOne.twoSum(numsFour, targetFour));
    }
}