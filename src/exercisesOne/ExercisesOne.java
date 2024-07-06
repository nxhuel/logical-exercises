package exercisesOne;

import java.util.Arrays;

public class ExercisesOne {
    public static int[] twoSum(int[] nums, int target) {

        int result = 0;
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                result = nums[i] + nums[j];
                if (result == target) {
                    return new int[]{i, j};
                }
            }
        }

        return new int[]{};
    }

    public static void main(String[] args) {
        int[] doubleSuma = twoSum(new int[]{2, 7, 11, 15}, 9);
        System.out.println(Arrays.toString(doubleSuma));
  }
}
